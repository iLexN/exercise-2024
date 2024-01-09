<?php

declare(strict_types=1);


namespace App\Controller;


use Hello\HiReply;
use Hello\HiUser;

class HiController
{
    public function sayHello(HiUser $user): HiReply
    {
        $message = new HiReply();
        $message->setMessage("Hello World");
        $message->setUser($user);
        return $message;
    }
}