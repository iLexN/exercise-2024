<?php

declare(strict_types=1);

use function Hyperf\Support\env;

return [
    'secret' => env('JWT_SECRET', 'please-fill-in'),
    'issuer' => env('JWT_ISSUER', 'please-fill-in'),
];