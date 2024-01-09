<?php

declare(strict_types=1);


namespace App\Controller;

use App\Services\Queue\Player\PlayerInfo;
use App\Services\Queue\Player\PlayerQueueService;
use Hyperf\Logger\LoggerFactory;

class PlayerQueue extends AbstractController
{
    private \Psr\Log\LoggerInterface $logger;

    public function __construct(
        private readonly PlayerQueueService $playerQueueService,
        LoggerFactory $loggerFactory,
    ) {
        $this->logger = $loggerFactory->get('log');
    }

    public function queue(): array
    {
        $player = PlayerInfo::fromRequest($this->request);
        $result = $this->playerQueueService->push($player);
        $this->logger->info("play queue result", [$result]);
        return [
            'result' => $result
        ];
    }
}