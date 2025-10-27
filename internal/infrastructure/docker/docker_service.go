package docker

import (
	"context"
	"fmt"
	"strings"

	appPor "github.com/jairoprogramador/fastdeploy/internal/application/ports"
	docPor "github.com/jairoprogramador/fastdeploy/internal/domain/docker/ports"
	docVos "github.com/jairoprogramador/fastdeploy/internal/domain/docker/vos"
	"github.com/jairoprogramador/fastdeploy/internal/infrastructure/executor"
)

type DockerService struct {
	exec   executor.CommandExecutor
	logger appPor.LogMessage
}

func NewDockerService(exec executor.CommandExecutor, logger appPor.LogMessage) docPor.DockerService {
	return &DockerService{
		exec:   exec,
		logger: logger,
	}
}

func (s *DockerService) Check(ctx context.Context) error {
	s.logger.Detail("Checking for Docker...")
	err := s.exec.Execute(ctx, "docker --version", "")
	if err != nil {
		return err
	}
	s.logger.Detail("Docker check successful")
	return nil
}

func (s *DockerService) Build(ctx context.Context, opts docVos.BuildOptions) error {
	s.logger.Detail(fmt.Sprintf("Building image: %s", opts.Image.FullName()))

	var commandBuilder strings.Builder
	commandBuilder.WriteString("docker build")

	for key, val := range opts.Args {
		commandBuilder.WriteString(fmt.Sprintf(" --build-arg %s=%s", key, val))
	}

	commandBuilder.WriteString(fmt.Sprintf(" -t %s", opts.Image.FullName()))

	if opts.Dockerfile != "" {
		commandBuilder.WriteString(fmt.Sprintf(" -f %s", opts.Dockerfile))
	}

	commandBuilder.WriteString(fmt.Sprintf(" %s", opts.Context))

	err := s.exec.Execute(ctx, commandBuilder.String(), opts.Context)
	if err != nil {
		return err
	}
	s.logger.Detail("Build image successful")
	return nil
}

func (s *DockerService) Run(ctx context.Context, opts docVos.RunOptions) error {
	s.logger.Detail("Running image...")
	var commandBuilder strings.Builder
	commandBuilder.WriteString("docker run")

	if opts.RemoveOnExit {
		commandBuilder.WriteString(" --rm")
	}

	for _, value := range opts.Groups {
		commandBuilder.WriteString(fmt.Sprintf(" --group-add %s", value))
	}

	if opts.Interactive {
		commandBuilder.WriteString(" -i")
	}
	if opts.AllocateTTY {
		commandBuilder.WriteString(" -t")
	}
	if opts.WorkDir != "" {
		commandBuilder.WriteString(fmt.Sprintf(" -w %s", opts.WorkDir))
	}

	for key, val := range opts.EnvVars {
		commandBuilder.WriteString(fmt.Sprintf(" -e %s=%s", key, val))
	}

	for _, vol := range opts.Volumes {
		commandBuilder.WriteString(fmt.Sprintf(" -v %s:%s", vol.HostPath, vol.ContainerPath))
	}

	commandBuilder.WriteString(fmt.Sprintf(" %s", opts.Image.FullName()))
	commandBuilder.WriteString(fmt.Sprintf(" %s", opts.Command))

	err := s.exec.Execute(ctx, commandBuilder.String(), "")
	if err != nil {
		return err
	}
	s.logger.Detail("Run image successful")
	return nil
}

   /*  docker run --rm \
      -e ARM_ACCESS_TOKEN="<pega-el-token-aqui>" \
      -e ARM_SUBSCRIPTION_ID="<tu-id-de-suscripcion>" \
      -e FASTDEPLOY_HOME=/home/fastdeploy/.fastdeploy \
      -v /home/jairo/.m2/:/home/fastdeploy/.m2 \
      -v /home/jairo/Developer/java/test:/home/fastdeploy/app \
      -v /home/jairo/Developer/go/dirFastDeploy:/home/fastdeploy/.fastdeploy \
      shikigami-mydeploy:latest test . */