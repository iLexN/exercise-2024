<?php

declare(strict_types=1);
/**
 * This file is part of Hyperf.
 *
 * @link     https://www.hyperf.io
 * @document https://hyperf.wiki
 * @contact  group@hyperf.io
 * @license  https://github.com/hyperf/hyperf/blob/master/LICENSE
 */

namespace App\Controller;


use App\Grpc\HiClient;
use App\Services\Jwt\JwtUserInfo;
use Hello\HiReply;
use Hello\HiUser;
use Hyperf\Logger\LoggerFactory;
use Psr\Log\LoggerInterface;

class IndexController extends AbstractController
{
    private HiClient $localClient;
    private HiClient $remoteClient;
    private HiClient $goClient;
    private LoggerInterface $logger;

    public function __construct(LoggerFactory $loggerFactory)
    {
        $this->remoteClient = new HiClient('gserver:9503', [
            'credentials' => null,
        ]);
        $this->localClient = new HiClient('127.0.0.1:9503', [
            'credentials' => null,
        ]);
        $this->goClient = new HiClient('go-grpc-server:50051', [
            'credentials' => null,
        ]);
        $this->logger = $loggerFactory->get('log');
    }

    public function index(): array
    {
        return [
            'message' => "Hello, World!",
        ];
    }

    public function hello(): array
    {
        // This client is coroutine-safe and can be reused
        $client = $this->localClient;

        return $this->say($client);
    }

    private function say(HiClient $client): array
    {
        $request = new HiUser();
        $request->setName('hyperf');
        $request->setSex(1);

        try {
            [$reply, $status] = $client->sayHello($request);
        } catch (\Throwable $exception) {
            return [
                'message' => $exception->getMessage(),
                'user_name' => '',
                'user_sex' => -1,
                'status' => -1,
            ];
        }

        $this->logger->info('replay', [$reply]);

        $message = $reply->getMessage();
        $user = $reply->getUser();

        return [
            'message' => $message,
            'user_name' => $user->getName(),
            'user_sex' => $user->getSex(),
            'status' => $status,
        ];
    }

    public function helloWorld(): array
    {
        // This client is coroutine-safe and can be reused
        $client = $this->remoteClient;

        return $this->say($client);
    }

    public function helloGo(): array
    {
        // seem call golang grpc few sec after have error
        // This client is coroutine-safe and can be reused
        //$client = $this->goClient;
        // after changed to create obj, error gone
        $client = new HiClient('go-grpc-server:50051', [
            'credentials' => null,
        ]);

        return $this->say($client);
    }

    public function tryAdd(): array
    {
        /** @var JwtUserInfo $user */
        $user = $this->request->getAttribute('user', null);

        return ['userInfo' => $user->toArray()];
    }
}
