<?php

// autoload_static.php @generated by Composer

namespace Composer\Autoload;

class ComposerStaticInitcb37612fede30e4cdc6a09c93266c223
{
    public static $files = array (
        '9e4824c5afbdc1482b6025ce3d4dfde8' => __DIR__ . '/..' . '/league/csv/src/functions_include.php',
    );

    public static $prefixLengthsPsr4 = array (
        'L' => 
        array (
            'League\\Csv\\' => 11,
        ),
        'I' => 
        array (
            'Ilexn\\1brc\\' => 11,
        ),
    );

    public static $prefixDirsPsr4 = array (
        'League\\Csv\\' => 
        array (
            0 => __DIR__ . '/..' . '/league/csv/src',
        ),
        'Ilexn\\1brc\\' => 
        array (
            0 => __DIR__ . '/../..' . '/src',
        ),
    );

    public static $classMap = array (
        'Composer\\InstalledVersions' => __DIR__ . '/..' . '/composer/InstalledVersions.php',
    );

    public static function getInitializer(ClassLoader $loader)
    {
        return \Closure::bind(function () use ($loader) {
            $loader->prefixLengthsPsr4 = ComposerStaticInitcb37612fede30e4cdc6a09c93266c223::$prefixLengthsPsr4;
            $loader->prefixDirsPsr4 = ComposerStaticInitcb37612fede30e4cdc6a09c93266c223::$prefixDirsPsr4;
            $loader->classMap = ComposerStaticInitcb37612fede30e4cdc6a09c93266c223::$classMap;

        }, null, ClassLoader::class);
    }
}
