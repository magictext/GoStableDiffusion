#!/bin/bash

# Init stable-diffusion.cpp
git submodule init
git submodule update

# Init stable-diffusion.cpp submodules
cd stable-diffusion.cpp
git submodule init
git submodule update

# Build static library
mkdir build
cd build

if [ "${CUDA:-0}" = "1" ]; then
    SD_CUDA="ON"
else
    SD_CUDA="OFF"
fi

if [ "${METAL:-0}" = "1" ]; then
    SD_METAL="ON"
else
    SD_METAL="OFF"
fi

if [ "${VULKAN:-0}" = "1" ]; then
    SD_VULKAN="ON"
else
    SD_VULKAN="OFF"
fi

cmake .. -DSD_CUDA=$SD_CUDA -DSD_METAL=$SD_METAL -DSD_VULKAN=$SD_VULKAN -DSD_SHARED_LIBRARY=ON

cmake --build . --config Release -j$(nproc)