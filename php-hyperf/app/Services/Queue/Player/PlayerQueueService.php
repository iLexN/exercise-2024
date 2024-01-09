<?php

declare(strict_types=1);


namespace App\Services\Queue\Player;

use Hyperf\AsyncQueue\Driver\DriverFactory;
use Hyperf\AsyncQueue\Driver\DriverInterface;

class PlayerQueueService
{
    private DriverInterface $driver;

    public function __construct(DriverFactory $driverFactory)
    {
        $this->driver = $driverFactory->get('default');
    }

    public function push(PlayerInfo $playerInfo){
        return $this->driver->push(new PlayerJob($playerInfo));
    }
}