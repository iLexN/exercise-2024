<?php

declare(strict_types=1);


namespace App\Services\Login;

use Hyperf\HttpServer\Contract\RequestInterface;

readonly class LoginForm
{
    private function __construct(
        public string $username,
        public string $password,
    )
    {
    }

    public static function fromRequest(RequestInterface $request): LoginForm
    {
        $body = $request->getParsedBody();

        return new self(
            $body['username'] ?? null,
            $body['password'] ?? null,
        );
    }
}