<?php

declare(strict_types=1);

namespace App\Kafka\Consumer;

use App\Services\Queue\Player\PlayerInfo;
use Hyperf\Kafka\AbstractConsumer;
use Hyperf\Kafka\Annotation\Consumer;
use Hyperf\Logger\LoggerFactory;
use longlang\phpkafka\Consumer\ConsumeMessage;

#[Consumer(topic: 'player', groupId: "hyperf2" ,autoCommit: true, nums: 2)]
class KafkaConsumer extends AbstractConsumer
{
    private \Psr\Log\LoggerInterface $logger;

    public function __construct(
        LoggerFactory $loggerFactory,
    )
    {
        $this->logger = $loggerFactory->get('kafak');
    }

    public function consume(ConsumeMessage $message)
    {
        //$message->getConsumer()->ack($message);

        //var_dump($message->getTopic().':'.$message->getKey().':'.$message->getValue());
        $this->logger->info('kafak-consumer', [
            'topic' => $message->getTopic(),
            'key' => $message->getKey(),
            'value' => (PlayerInfo::fromMessage($message))->toArray(),
        ]);

    }
}
