#include <GLFW/glfw3.h>

#include "../tools/definitions.h"

namespace Vuelto {

class Window;

void Init();
void Terminate();

namespace Application {

Window CreateWindow(int width, int height, const char *title);
void DestroyWindow(Window win);

}  // namespace Application

class Window {
 public:
  GLFWwindow *window;
  int width;
  int height;
  const char *title;
  bool WindowShouldClose();
  void WindowRefresh();
};

}  // namespace Vuelto
