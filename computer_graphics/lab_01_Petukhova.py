import pyglet
from pyglet.gl import *
from pyglet.window import mouse

import random


window = pyglet.window.Window()

global array
array = [0, 0, 0]
global array1
array1 = [50, 50, 300, 100, 100, 200]
@window.event

def on_draw():
    glClear(GL_COLOR_BUFFER_BIT);

    glBegin(GL_TRIANGLES);


    glColor3f(array[0], array[1], array[2]);

    glVertex2f(array1[0], array1[1]);
    glVertex2f(array1[2], array1[3]);
    glVertex2f(array1[4], array1[5]);
    glEnd()


@window.event
def on_mouse_press(x, y, button, modifiers):
    if button == mouse.LEFT:
        array[0] += random.random()
        array[1] = random.random()
        array[2] = random.random()
        array1[0] += 30
        array1[1] += 30
        array1[2] += 30
        array1[3] += 30
        array1[4] += 30
        array1[5] += 30


if __name__ == '__main__':
    pyglet.app.run()
