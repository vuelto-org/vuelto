#include "shader.hpp"

#include <cstring>
#include <fstream>
#include <sstream>

namespace Vuelto {
namespace Shader {

const char* ReadShaderFile(const char* filePath) {
  std::ifstream file(filePath);
  std::stringstream buffer;
  buffer << file.rdbuf();

  // Store the content in a string and return its c_str()
  std::string content = buffer.str();
  char* shaderSource = new char[content.size() + 1];
  std::strcpy(shaderSource, content.c_str());

  return shaderSource;
}

}  // namespace Shader
}  // namespace Vuelto
