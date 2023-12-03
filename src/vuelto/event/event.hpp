#pragma once

#include "../tools/definitions.hpp"
#include "../core/app.hpp"

namespace Vuelto {
namespace Events {

void PollEvents();

}

namespace Input {

bool IsKeyPressed(Vuelto::Window window, int keyCode);
bool IsKeyPressedOnce(Vuelto::Window window, int key);

}

}  // namespace Vuelto
