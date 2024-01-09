<?php

declare(strict_types=1);

namespace App\Services\Queue\Player;

use Hyperf\HttpServer\Contract\RequestInterface;

readonly class PlayerInfo
{
    public function __construct(
        public string $name,
        public int $age,
        public string $sex,
    ) {
    }

    public static function fromRequest(RequestInterface $request): PlayerInfo
    {
        $body = $request->getParsedBody();
        return new self(
            $body['name'] ?? null,
            $body['age'] ?? null,
            $body['sex'] ?? null,
        );
    }
}