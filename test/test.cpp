#include "../src/Festi.h"

int main() {
  FS_Init();

  FS_Window* window = FS_CreateWindow(800, 600, "Test");

  while (!FS_WindowShouldClose()) {
    FS_SetBackgroundColor(0.5, 0.5, 0.5);

    FS_DrawRect(0, 0, 0.5, 0.5, 0.1, 0.1, 0.1);

    FS_Refresh();
  }
  return 0;
}
