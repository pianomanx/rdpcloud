<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: models/fileio/quota_entry.proto

namespace GPBMetadata\Models\Fileio;

class QuotaEntry
{
    public static $is_initialized = false;

    public static function initOnce() {
        $pool = \Google\Protobuf\Internal\DescriptorPool::getGeneratedPool();

        if (static::$is_initialized == true) {
          return;
        }
        $pool->internalAddGeneratedFile(
            '
�
models/fileio/quota_entry.proto

QuotaEntry
username (	
quota_threshold (
quota_limit (

quota_used (
account_status (
QuotaEntry_6
quota_threshold (
quota_limit ("i

quota_threshold (
quota_limit (

quota_used (
account_status (
        , true);

        static::$is_initialized = true;
    }
}
