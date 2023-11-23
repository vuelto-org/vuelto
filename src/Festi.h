#ifdef FESTI_H
#include "festi/window.c"
#endif

#include <GLFW/glfw3.h>

typedef struct {
	GLFWwindow* window;
	int width;
	int height;
	const char* title;
} FS_Window;

// Window.c
void FS_Init();
FS_Window* FS_CreateWindow(int width, int height, const char* title);
int FS_WindowShouldClose();
void FS_Refresh();

// Draw.c
void FS_DrawRect(float x, float y, float height, float width, const float color[3]);
void FS_SetBackgroundColor(const float color[3]);