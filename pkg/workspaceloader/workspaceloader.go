package workspaceloader

import (
	"fmt"

	"go.i3wm.org/i3/v4"
)

func LoadWorkspace() {
	fmt.Println("load-workspace")
}

func init1(debug bool) {
	//    self.debug = debug
	//
	//    self.workspaces = {
	//        1: ['alacritty', 'firefox'],
	//        2: ['firefox', 'firefox', 'alacritty'],
	//        3: ['virtualbox'],
	//        4: ['slack', 'teams', 'firefox'],
	//        8: ['alacritty'],
	//        9: ['alacritty', 'alacritty', 'alacritty']
	//
	//    }
	//
	//    self.applications = {
	//        'alacritty': 'alacritty --working-directory ~',
	//        'firefox': 'firefox',
	//        'virtualbox': 'vboxmanage startvm \'Windows10\'',
	//        'slack': 'slack',
	//        'teams': 'teams'
	//    }
}

func run(workspace i3.Workspace) {
	//    commands = []
	//    has_errors = False
	//
	//    for workspace in workspaces:
	//        try:
	//            commands.append('i3-msg "{0}"'.format(self.build_command(workspace)))
	//        except Exception as error:
	//            has_errors = True
	//            print(error)
	//
	//    self.print_debug(commands)
	//    self.run_commands(commands, has_errors)
}

func run_commands() {
	//    if has_errors:
	//        return
	//
	//    if self.debug:
	//        return
	//
	//    for cmd in commands:
	//        subprocess.call(cmd, shell=True, stdout=subprocess.DEVNULL)
}

func build_command() bool {
	//    if workspace not in self.workspaces:
	//        raise Exception('Workspace "{0}" not configured'.format(workspace))
	//
	//    cmd = self.build_workspace_cmd(workspace)
	//
	//    for application in self.workspaces[workspace]:
	//        cmd += self.build_application_cmd(application)
	//
	//    return cmd.rstrip()
	return false
}

func build_workspace_cmd() string {
	//    workspace_cmd = 'workspace {0}:{0};'.format(workspace)
	//    layout_cmd = 'append_layout ~/.config/i3/workspaces/workspace-{0}.json; '.format(workspace)
	//    return '{0} {1}'.format(workspace_cmd, layout_cmd)
	return ""
}

func build_application_cmd() string {
	//    cmd = self.applications[application]
	//    return 'exec --no-startup-id {0}; '.format(cmd)
	return ""
}
