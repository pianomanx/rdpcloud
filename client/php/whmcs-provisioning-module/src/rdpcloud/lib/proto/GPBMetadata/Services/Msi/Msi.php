<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: services/msi/msi.proto

namespace GPBMetadata\Services\Msi;

class Msi
{
    public static $is_initialized = false;

    public static function initOnce() {
        $pool = \Google\Protobuf\Internal\DescriptorPool::getGeneratedPool();

        if (static::$is_initialized == true) {
          return;
        }
        \GPBMetadata\Services\Msi\SystemStatus::initOnce();
        $pool->internalAddGeneratedFile(
            '
�
services/msi/msi.protoservices.msi2[
MsiT
GetProducts .services.msi.GetProductsRequest!.services.msi.GetProductsResponse" B1Z/github.com/s77rt/rdpcloud/proto/go/services/msibproto3'
        , true);

        static::$is_initialized = true;
    }
}

