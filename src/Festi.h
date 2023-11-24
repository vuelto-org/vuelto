#ifdef __cplusplus
extern "C" {
#endif

#include <GLFW/glfw3.h>

typedef struct {
  GLFWwindow *window;
  int width;
  int height;
  const char *title;
} FS_Window;

// Window.c
void FS_Init();
FS_Window *FS_CreateWindow(int width, int height, const char *title);
int FS_WindowShouldClose();
void FS_Refresh();

// Draw.c
void FS_DrawRect(float x, float y, float height, float width, float color1, float color2, float color3);
void FS_SetBackgroundColor(float color1, float color2, float color3);

#ifdef __cplusplus
}
#endif
