<?php

declare(strict_types=1);

namespace App\Services\Queue\Player;

use Hyperf\AsyncQueue\Job;

class PlayerJob extends Job
{

    public function __construct(
        private readonly PlayerInfo $playerInfo,
    ) {
    }

    public function handle(): void
    {
        var_dump($this->playerInfo);
    }
}