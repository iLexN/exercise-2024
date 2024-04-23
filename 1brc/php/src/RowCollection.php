<?php

declare(strict_types=1);


namespace Ilexn\OneBrc;

final readonly class RowCollection
{
    private array $citys;

    public function __construct()
    {
        $this->citys = [];
    }

    public function add(string $row): void
    {
        $data = RowData::create($row);

        if (!$this->isExist($data)) {

        }

    }

    private function isExist(RowData $city): bool
    {
        return isset($this->citys[$city->name]);
    }

}