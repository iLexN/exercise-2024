<?php

declare(strict_types=1);


namespace App\Services\Jwt;

readonly class JwtUserInfo
{
    private function __construct(
        public string $role,
        public string $iss,
        public string $sub,
        public int $exp,
        public int $nbf,
        public int $iat,
    ) {
    }

    public function toArray(): array
    {
        return [
            'role' => $this->role,
            'iss' => $this->iss,
            'sub' => $this->sub,
            'exp' => $this->exp,
            'nbf' => $this->nbf,
            'iat' => $this->iat,
        ];
    }

    public static function createFromStd(\stdClass $info): JwtUserInfo
    {
        return new self(
            $info->role,
            $info->iss,
            $info->sub,
            $info->exp,
            $info->nbf,
            $info->iat,
        );
    }
}