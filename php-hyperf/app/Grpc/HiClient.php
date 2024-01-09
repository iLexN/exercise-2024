<?php

declare(strict_types=1);


namespace App\Grpc;

use Hello\HiReply;
use Hello\HiUser;
use Hyperf\GrpcClient\BaseClient;

class HiClient extends BaseClient
{

    public function sayHello(HiUser $argument): array
    {
        return $this->_simpleRequest(
            '/hello.hi/sayHello',
            $argument,
            [HiReply::class, 'decode']
        );
    }
}