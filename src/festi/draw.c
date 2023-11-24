#include "definitions.h"

void FS_DrawRect(float x, float y, float height, float width, float color1, float color2, float color3) {
  glBegin(GL_QUADS);
  glColor3f(color1, color2, color3);
  glVertex2f(x, y);
  glVertex2f(x + width, y);
  glVertex2f(x + width, y + height);
  glVertex2f(x, y + height);
  glEnd();
}

// clang-format off
void FS_SetBackgroundColor(float color1, float color2, float color3) {
    glClearColor(color1, color2, color3, 1);
}
// clang-format on
