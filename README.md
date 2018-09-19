# pencere

IT IS JUST A PROTOTYPE FOR NOW AND SUBJECT TO HEAVYLY CHANGE 

A new Terminal UI kit for go

<img  width="480" height="473" src="https://github.com/ilkeraksu/pencere/blob/master/media/pencere_preview_01.gif">

### What is new
#### Full rendering with windows (rectangles)
Every ui element is a window (pencere) and ui is built by compositing new windows adding as child window. Every window can focus, receive mouse events, keyboard events, drag, drop, render and so on. 

Pencere takes and dispatches terminal events to windows and  renders windows based on Zindexes and window attributes like Visible and HasBorder  

#### Layout
Every window has it's Layout event to change layout of itself or childs. Also every window has LayoutOrder attribute to take precedence on layout event order. So we can easly implement window slider where slider manages side windows layouts 

#### Drag and Drop
Windows support Draggable,Droppable attributes to be come eligible for Drag and Drop events. OnDragBegin, OnDragEnd, OnDragging, OnDragEnter, OnDragOver, OnDragLeave, CanDrop and OnDrop events are supported.

#### Focusing
Windows can receive keyboard events when they are focused. Managed by CanFocus attribute. 

#### Mouse events
Windows suport mouse events and events canbe received if relavent eventhandler is implemented, if not event may be handled by parent windows

#### New Event system (as message bus)
New event system will be implemented as message bus to support inter window communication

#### Sidebars
Windows has left right sidebar , top and bottom toolbars as window (pencere) means can be implemented as menu, toolbar or sticked toolboxes

### Menu system 
Menu systems and menu events ( in progress)


Because usage of window ui is very elastic and mature concept, many many widgets and features can be implemented.


#### Used libraries
Inspired and used some code from these libraries. Basicly tried to be like and compatible to reuse of widgets and other stuff 
-  https://github.com/nsf/termbox
   - Pencere use termbox for base Terminal functions like rendering and events

- [https://github.com/gizak/termui](https://github.com/gizak/termui) 
  - Buffer struct taken so widgets are almost compatible with minimal modifications 

- [https://github.com/cjbassi/termui](https://github.com/cjbassi/termui)
  - Some widgets taken and converted
  - Theming is inspired and taken (it subject to improve)

- https://github.com/marcusolsson/tui-go
  - Widgets are not dropin compatible but can easly be converted.
  - Runebuffer taken
  

