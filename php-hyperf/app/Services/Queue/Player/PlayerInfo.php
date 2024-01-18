<?php

declare(strict_types=1);

namespace App\Services\Queue\Player;

use Hyperf\HttpServer\Contract\RequestInterface;
use longlang\phpkafka\Consumer\ConsumeMessage;

readonly class PlayerInfo
{
    public function __construct(
        public string $name,
        public int $age,
        public string $sex,
    ) {
    }

    /**
     * @throws \JsonException
     */
    public function toJson(): string
    {
        return \json_encode($this->toArray(), JSON_THROW_ON_ERROR);
    }

    public function toArray(): array
    {
        return [
            'name' => $this->name,
            'age' => $this->age,
            'sex' => $this->sex,
        ];
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

    public static function fromMessage(ConsumeMessage $message)
    {
        $var = \json_decode($message->getValue(), true);
        return new self(
            $var['name'] ?? null,
            $var['age'] ?? null,
            $var['sex'] ?? null,
        );
    }

}