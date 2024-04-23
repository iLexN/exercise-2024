<?php

declare(strict_types=1);


namespace Ilexn\OneBrc;

final readonly class RowData
{
    public function __construct(
        public string $name,
        public float $measurement,
    ) {
    }

    public static function create(string $data): RowData
    {
        $city = explode(',', $data);
        return new self($city[0], (float)$city[1]);
    }
}