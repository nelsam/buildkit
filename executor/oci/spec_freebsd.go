package oci

import (
	"github.com/containerd/containerd/oci"
	"github.com/docker/docker/pkg/idtools"
	"github.com/moby/buildkit/solver/pb"
	specs "github.com/opencontainers/runtime-spec/specs-go"
	"github.com/pkg/errors"
)

func generateMountOpts(resolvConf, hostsFile string) ([]oci.SpecOpts, error) {
	return nil, nil
}

// generateSecurityOpts may affect mounts, so must be called after generateMountOpts
func generateSecurityOpts(mode pb.SecurityMode, apparmorProfile string, selinuxB bool) ([]oci.SpecOpts, error) {
	if mode == pb.SecurityMode_INSECURE {
		return nil, errors.New("no support for running in insecure mode on FreeBSD")
	}
	return nil, nil
}

// generateProcessModeOpts may affect mounts, so must be called after generateMountOpts
func generateProcessModeOpts(mode ProcessMode) ([]oci.SpecOpts, error) {
	if mode == NoProcessSandbox {
		return nil, errors.New("no support for NoProcessSandbox on FreeBSD")
	}
	return nil, nil
}

func generateIDmapOpts(idmap *idtools.IdentityMapping) ([]oci.SpecOpts, error) {
	if idmap == nil {
		return nil, nil
	}
	return nil, errors.New("no support for IdentityMapping on FreeBSD")
}

func generateRlimitOpts(ulimits []*pb.Ulimit) ([]oci.SpecOpts, error) {
	if len(ulimits) == 0 {
		return nil, nil
	}
	return nil, errors.New("no support for POSIXRlimit on FreeBSD")
}

// tracing is not implemented on FreeBSD
func getTracingSocketMount(socket string) *specs.Mount {
	return nil
}

// tracing is not implemented on FreeBSD
func getTracingSocket() string {
	return ""
}

func cgroupNamespaceSupported() bool {
	return false
}
