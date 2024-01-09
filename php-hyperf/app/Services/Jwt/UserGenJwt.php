<?php

declare(strict_types=1);


namespace App\Services\Jwt;

use App\Services\Login\LoginForm;

readonly class UserGenJwt
{
    private function __construct(
        public string $username,
        public string $role,
    )
    {
    }

    public static function createFromLoginForm(LoginForm $loginForm): UserGenJwt
    {
        return new self($loginForm->username, 'Admin');
    }
}