<?php

declare(strict_types=1);

namespace App\Middleware;

use Hyperf\Logger\LoggerFactory;
use Psr\Http\Message\ResponseInterface;
use Psr\Http\Server\MiddlewareInterface;
use Psr\Http\Message\ServerRequestInterface;
use Psr\Http\Server\RequestHandlerInterface;

class Performance implements MiddlewareInterface
{
    private \Psr\Log\LoggerInterface $logger;

    public function __construct(LoggerFactory $loggerFactory)
    {
        // The first parameter corresponds to the name of the log, and the second parameter corresponds to the key in config/autoload/logger.php
        $this->logger = $loggerFactory->get('log', 'default');
    }

    public function process(ServerRequestInterface $request, RequestHandlerInterface $handler): ResponseInterface
    {
        $start = microtime(true);

        $next = $handler->handle($request);

        $end = microtime(true);

        $this->log($start, $end, $request, $next);

        return $next;
    }

    private function log(float $start, float $end, ServerRequestInterface $request, ResponseInterface $response): void
    {
        $timeDiffMs = round(($end - $start) * 1000000,2);

//        var_dump("Time difference: ".round($timeDiffMs, 2)." µs");
        $this->logger->info("Request info", [
            'status'=> $response->getStatusCode(),
            'time' => $timeDiffMs . ' µs',
            'uri' => $request->getUri()->getPath(),
        ]);
    }
}
