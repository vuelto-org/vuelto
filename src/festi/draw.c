#include "definitions.h"

void FS_DrawRect(float x, float y, float height, float width,
                 const float color[3]) {
  glBegin(GL_QUADS);
  glColor3f(color[0], color[1], color[2]);
  glVertex2f(x, y);
  glVertex2f(x + width, y);
  glVertex2f(x + width, y + height);
  glVertex2f(x, y + height);
  glEnd();
}

void FS_SetBackgroundColor(const float color[3]) {
  glClearColor(color[0], color[1], color[2], 1);
}
