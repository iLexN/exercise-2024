<?php

declare(strict_types=1);

namespace App\Services\Jwt;

use Firebase\JWT\JWT;
use Firebase\JWT\Key;

class JwtService
{
    private string $secret;
    private string $issuer;

    public function __construct(
        \Hyperf\Contract\ConfigInterface $config,
    )
    {
        $this->secret = $config->get('jwt.secret');
        $this->issuer = $config->get('jwt.issuer');
    }

    public function genToken(UserGenJwt $userGenJwt): string
    {
        $key = $this->secret;
        $exp = time() + (24 * 60 * 60);
        $now = time();

        $payload = [
            'role' => $userGenJwt->role,
            'iss' => $this->issuer,
            'sub' => $userGenJwt->username,
            'iat' => $now,
            'exp' => $exp,
            'nbf' => $now,
        ];
        return JWT::encode($payload, $key, 'HS512');
    }

    public function decodeToken(string $token): JwtUserInfo
    {
        $key = $this->secret;
        $headers = new \stdClass();
        $decoded = JWT::decode($token, new Key($key, 'HS512'), $headers);
        return JwtUserInfo::createFromStd($decoded);
    }
}