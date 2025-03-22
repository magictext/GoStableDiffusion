@echo off

git submodule init
git submodule update

cd stable-diffusion.cpp
git submodule init
git submodule update

mkdir build
cd build

if "%CUDA%"=="1" (
    set SD_CUDA=ON
) else (
    set SD_CUDA=OFF
)
if "%METAL%"=="1" (
    set SD_METAL=ON
) else (
    set SD_METAL=OFF
)
if "%VULKAN%"=="1" (
    set SD_VULKAN=ON
) else (
    set SD_VULKAN=OFF
)

cmake .. -DSD_CUDA=%SD_CUDA% -DSD_METAL=%SD_METAL% -DSD_VULKAN=%SD_VULKAN% -DSD_BUILD_SHARED_LIBS=ON

cmake --build . --config Release -j%NUMBER_OF_PROCESSORS%
dir /s /b *.dll
copy bin\Release\libstable-diffusion.dll ..\..