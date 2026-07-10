In Doom Emacs, projects and workspaces work seamlessly together to isolate your code, open buffers, and window layouts per project. Projects (powered by Projectile) group files belonging to a repository, while Workspaces (powered by persp-mode) act as isolated virtual desktops that filter your buffer list to only show files relevant to that specific context. 

Here is a comprehensive breakdown of how to use and manage them effectively. 
Managing Projects (SPC p) 
Doom Emacs uses Projectile to automatically detect project roots if a directory contains a .git folder, a .project file, or standard build configuration files (like package.json or Cargo.toml). 

• SPC p p  – Switch project. This lists your known projects and prompts you to open one. 
• SPC p f – Find file in the current project. 
• SPC p a – Add a new project directory to Doom's known index manually. 
• SPC p s s – Search text across the entire project using Ripgrep. 
• SPC p r – View recently visited files only within the current project. [1, 5, 6, 7, 8]  

Managing Workspaces (SPC TAB) 
Workspaces allow you to compartmentalize multiple open projects. When you switch to a workspace, your buffer list () only displays the buffers loaded inside that workspace, avoiding messy multi-project clutter. 

• SPC TAB n – Create a new blank workspace. 
• SPC TAB r – Rename the current workspace. 
• SPC TAB d – Delete / Kill the current workspace. 
• SPC TAB . – Switch workspace by selecting from a list. 
• SPC TAB [ or ] – Cycle to the previous or next workspace. 
• SPC TAB 1-9 – Jump directly to a specific workspace number.

The Ultimate Pro-Workflow 
Instead of creating a workspace and then manually finding a project, combine them using the ultimate Doom Emacs shortcut: 

1. Press SPC p p to switch projects. 
2. Select your project. 
3. Doom will automatically spin up a brand new workspace named after that project, loading its files and completely isolating its buffer list. [1, 2, 14, 15]  

If you want to customize your workflow configuration or learn more about the workspace core settings, you can check out the official Doom Emacs UI Workspaces Documentation. 
Would you like to know how to configure Doom to automatically save your workspace sessions when you quit, or do you need help excluding specific folders from your project searches? 

AI responses may include mistakes.

[1] https://www.youtube.com/watch?v=Rx3wRl5d-J0
[2] https://practical.li/doom-emacs/basics/workspaces/
[3] https://www.reddit.com/r/emacs/comments/1or88u7/what_to_do_about_workspaces/
[4] https://noelwelsh.com/posts/doom-emacs/
[5] https://practical.li/doom-emacs/basics/projects/
[6] https://www.reddit.com/r/emacs/comments/6lgwme/how_do_you_handle_multiple_projectsworkspaces/
[7] https://0x7ffc.github.io/2019/my-emacs-journey/
[8] https://marketplace.visualstudio.com/items?itemName=Bearylabs.doom
[9] https://www.reddit.com/r/emacs/comments/11od5jk/two_projects_side_by_side/
[10] https://docs.doomemacs.org/v21.12/modules/ui/workspaces/
[11] https://valerioviperino.me/doom-emacs-handbook/
[12] https://github.com/doomemacs/core/issues/4117
[13] https://gist.github.com/darrylhebbes/5b68268b8a12053a57abf9b02f0ec61f
[14] https://github.com/doomemacs/doomemacs/issues/836
[15] https://www.matfournier.com/2019-10-05-scalaemacs/
[16] https://github.com/rcallaby/Emacs-Guide/blob/main/Lessons/Part-09-Doom-Emacs/doomemacs.md

