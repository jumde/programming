#include <iostream>
#include <cstring>
#include <memory>

std::unique_ptr<char[]> process_buffer(const char* input_buffer, size_t input_len, size_t& output_len) {
    if (input_buffer == NULL) {
        throw std::invalid_argument("Input buffer is NULL");
    }

    const size_t MAX_BUFFER_SIZE = 1024;
    if (input_len > MAX_BUFFER_SIZE) {
        throw std::length_error("Input buffer exceeds length");
    }

    std::unique_ptr<char[]> output_buffer(new char[input_len+1]);
    std::memcpy(output_buffer.get(), input_buffer, input_len);
    output_buffer[input_len] = '\0';

    for (size_t i=0; i < input_len; i++) {
        if (output_buffer[i] >= 'a' && output_buffer[i] <= 'z') {
            output_buffer[i] = output_buffer[i] - 32;
        }
    }

    output_len = input_len;
    return output_buffer;
}

int main() {
    try {
        const char* input = "Hello, World!";
        size_t input_len = strlen(input);
        size_t output_len = 0;

        auto output = process_buffer(input, input_len, output_len);
        std::cout << "Processed Buffer: " << output.get() << std::endl;
    }  catch (const std::exception &e) {
        std::cerr << "Error: " << e.what () << std::endl;
        return 1;
    }
    return 0;
}
