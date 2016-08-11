# i3title

Show title of focused window in i3wm with support of multiple outputs using i3's IPC (with [sometimesfood's](https://github.com/sometimesfood/i3ipc) fork of [i3ipc](https://github.com/proxypoke/i3ipc) library which is a little bit more up to date). 

## Usage

`i3title`
or
`i3title -o <OUTPUT>`

To get available outputs you may use oneliner below  
`xrandr | grep " connected" | cut -d' ' -f1`
