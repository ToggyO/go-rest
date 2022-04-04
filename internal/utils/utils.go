package utils

import (
    "fmt"
    "os"
    "path/filepath"
)

func GetEnv(key, fallback string) string {
    if value, ok := os.LookupEnv(key); ok {
        return value
    }
    return fallback
}

func GetEnvFilePathFromRoot(appEnvVariableName, fallback string) string {
    return GetEnvFilePath("./", appEnvVariableName, fallback)
}

func GetEnvFilePath(containingFolderPath, appEnvVariableName, fallback string) string {
    goEnv := GetEnv(appEnvVariableName, fallback)
    pwd, _ := os.Getwd()
    return filepath.Join(pwd, fmt.Sprintf("%s.env.%s", containingFolderPath, goEnv))
}
