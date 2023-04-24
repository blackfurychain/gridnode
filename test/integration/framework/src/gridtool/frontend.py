from gridtool import command, project
from gridtool.common import *

GIT_REPO_PROXIES = "git@github.com:Gridironchain/proxies.git"
GITHUB_REPO_UI = "https://github.com/Gridironchain/gridchain-ui.git"


# Vercel-based proxy to "hide" all URLs behind a single proxy due to CORS policy.
# You need an account, configuration is in ~/.cache/com.vercel/cli and ~/.local/share/com.vercel.cli.
class ProxiesProject:
    def __init__(self, cmd: command.Command, dir: str):
        self.cmd = cmd
        self.dir = dir


class FrontendProject:
    def __init__(self, cmd: command.Command, dir: str):
        self.cmd = cmd
        self.dir = dir

    def v2_configure_for_localnet_host(self, host: str, api_port: int, rpc_port: str):
        pass


def run_local_ui(cmd: command.Command):
    prj = project.Project(cmd, project_dir())
    working_dir = project_dir("test", "integration", "framework", "build", "repos")
    cmd.mkdir(working_dir)
    proxies_repo_dir = os.path.join(working_dir, "proxies")
    ui_repo_dir = os.path.join(working_dir, "gridchain-ui")
    cmd.rmdir(proxies_repo_dir)
    cmd.rmdir(ui_repo_dir)
    prj.git_clone(GIT_REPO_PROXIES, proxies_repo_dir)
    prj.git_clone(GITHUB_REPO_UI, ui_repo_dir)
