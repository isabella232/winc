package sandbox

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"

	"code.cloudfoundry.org/winc/hcsclient"

	"github.com/Microsoft/hcsshim"
)

var sandboxFiles = []string{"Hives", "initialized", "sandbox.vhdx", "layerchain.json"}

//go:generate counterfeiter . SandboxManager
type SandboxManager interface {
	Create(rootfs string) (string, error)
	Delete() error
	BundlePath() string
	Mount(pid int) error
	Unmount(pid int) error
}

//go:generate counterfeiter . Mounter
type Mounter interface {
	SetPoint(string, string) error
	DeletePoint(string) error
}

type sandboxManager struct {
	hcsClient  hcsclient.Client
	id         string
	driverInfo hcsshim.DriverInfo
	mounter    Mounter
	volumePath string
}

func NewManager(hcsClient hcsclient.Client, mounter Mounter, depotDir, containerId string) SandboxManager {
	driverInfo := hcsshim.DriverInfo{
		HomeDir: depotDir,
		Flavour: 1,
	}

	return &sandboxManager{
		hcsClient:  hcsClient,
		mounter:    mounter,
		id:         containerId,
		driverInfo: driverInfo,
	}
}

func (s *sandboxManager) Create(rootfs string) (string, error) {
	err := os.MkdirAll(s.driverInfo.HomeDir, 0755)
	if err != nil {
		return "", err
	}

	_, err = os.Stat(rootfs)
	if os.IsNotExist(err) {
		return "", &MissingRootfsError{Msg: rootfs}
	} else if err != nil {
		return "", err
	}

	parentLayerChain, err := ioutil.ReadFile(filepath.Join(rootfs, "layerchain.json"))
	if err != nil {
		return "", &MissingRootfsLayerChainError{Msg: rootfs}
	}

	parentLayers := []string{}
	if err := json.Unmarshal(parentLayerChain, &parentLayers); err != nil {
		return "", &InvalidRootfsLayerChainError{Msg: rootfs}
	}
	sandboxLayers := append([]string{rootfs}, parentLayers...)

	if err := s.hcsClient.CreateSandboxLayer(s.driverInfo, s.id, rootfs, sandboxLayers); err != nil {
		return "", err
	}

	if err := s.hcsClient.ActivateLayer(s.driverInfo, s.id); err != nil {
		return "", err
	}

	if err := s.hcsClient.PrepareLayer(s.driverInfo, s.id, sandboxLayers); err != nil {
		return "", err
	}

	sandboxLayerChain, err := json.Marshal(sandboxLayers)
	if err != nil {
		return "", err
	}

	err = os.MkdirAll(filepath.Join(s.driverInfo.HomeDir, s.id), 0755)
	if err != nil {
		return "", err
	}

	if err := ioutil.WriteFile(filepath.Join(s.driverInfo.HomeDir, s.id, "layerchain.json"), sandboxLayerChain, 0755); err != nil {
		return "", err
	}

	volumePath, err := s.hcsClient.GetLayerMountPath(s.driverInfo, s.id)
	if err != nil {
		return "", err
	} else if volumePath == "" {
		return "", &hcsclient.MissingVolumePathError{Id: s.id}
	}

	s.volumePath = volumePath

	return volumePath, nil
}

func (s *sandboxManager) Delete() error {
	if err := s.hcsClient.UnprepareLayer(s.driverInfo, s.id); err != nil {
		return err
	}

	if err := s.hcsClient.DeactivateLayer(s.driverInfo, s.id); err != nil {
		return err
	}

	for _, f := range sandboxFiles {
		layerFile := filepath.Join(s.driverInfo.HomeDir, f)
		if err := os.RemoveAll(layerFile); err != nil {
			return &UnableToDestroyLayerError{Msg: layerFile}
		}
	}

	return nil
}

func (s *sandboxManager) BundlePath() string {
	return filepath.Join(s.driverInfo.HomeDir, s.id)
}

func (s *sandboxManager) mountPath(pid int) string {
	return filepath.Join("c:\\", "proc", strconv.Itoa(pid))
}

func (s *sandboxManager) rootPath(pid int) string {
	return filepath.Join(s.mountPath(pid), "root")
}

func (s *sandboxManager) Mount(pid int) error {
	if err := os.MkdirAll(s.rootPath(pid), 0755); err != nil {
		return err
	}

	return s.mounter.SetPoint(s.rootPath(pid), s.volumePath)
}

func (s *sandboxManager) Unmount(pid int) error {
	if err := s.mounter.DeletePoint(s.rootPath(pid)); err != nil {
		return err
	}

	return os.RemoveAll(s.mountPath(pid))
}
