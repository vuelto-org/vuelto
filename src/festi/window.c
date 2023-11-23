#include "definitions.h"

FS_Window* win;

void FS_Init() {
	if (!glfwInit())
        printf("GLFW Init failed\n");
}

FS_Window* FS_CreateWindow(int width, int height, const char* title) {
	GLFWwindow* glfw_window = glfwCreateWindow(width, height, title, NULL, NULL);
	if (!glfw_window)
	{
		glfwTerminate();
		printf("GLFW Window creation failed\n");
	}

	glfwMakeContextCurrent(glfw_window);

	FS_Window* window;
	window->width = width;
	window->height = height;
	window->title = title;
	window->window = glfw_window;

	win = window;

	return window;
}

int FS_WindowShouldClose() {
	return glfwWindowShouldClose(win->window);
}

void FS_Refresh() {
	glfwSwapBuffers(win->window);
	glfwPollEvents();
    glClear(GL_COLOR_BUFFER_BIT);
}


