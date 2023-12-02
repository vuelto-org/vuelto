#include "../include/vuelto/Vuelto.hpp"

int main() {
  Vuelto::Application::Init();

  Vuelto::Window win = Vuelto::Application::CreateWindow(800, 600, "test", true);

  Vuelto::Renderer2D renderer = Vuelto::Application::CreateRenderer2D(win);

  while (!win.WindowShouldClose()) {
    renderer.DrawRect(0.1, 0.1, 0.5, 0.5, 0.6, 0.4, 0.8);
    win.Refresh();
  }

  Vuelto::Application::Terminate();
  return 0;
}
