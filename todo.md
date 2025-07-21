In change the makefile to be like the makefile in the gerasimova businesscard.
==================
In the base template directory the styles and scripts for both are in the styles/scripts of the base instead of separated in the footer and navbar directories.

The js/css for all home components are in the base styles/scripts files instead of in the directories of each component.

The individual js/css files for each of the components are not passed to the home.templ file. They need to have a similar funciton to the homeComponents one and be concatenated with the global styles/scripts of the home page.

Each style/script of each of the components should not be wrapped with opening and closing tags. Instead there should be just one set of global opening and closing tags.

Figure out which styles/scripts shoud be global and which should not.
===========

Figure out how to do languages.
