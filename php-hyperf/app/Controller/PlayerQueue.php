<?php

declare(strict_types=1);


namespace App\Controller;

use App\Services\Queue\Player\PlayerInfo;
use App\Services\Queue\Player\PlayerQueueService;
use Hyperf\Logger\LoggerFactory;
use Hyperf\Kafka\Producer;
use Psr\Log\LoggerInterface;

class PlayerQueue extends AbstractController
{
    private LoggerInterface $logger;

    public function __construct(
        private readonly PlayerQueueService $playerQueueService,
        LoggerFactory $loggerFactory,
        private  readonly Producer $producer
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

    /**
     * @throws \JsonException
     */
    public function kafkaQueue(): array
    {
        $player = PlayerInfo::fromRequest($this->request);
        $this->producer->send(
            'player',
            $player->toJson(),
            'A-' . \random_int(1,999)
        );


        return [
            'result' => 'done',
            'player' => $player->toArray(),
        ];
    }
}