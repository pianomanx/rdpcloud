<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: services/fileio/disk_management.proto

namespace Services\Fileio;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>services.fileio.SetDefaultQuotaRequest</code>
 */
class SetDefaultQuotaRequest extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>string volume_path = 1;</code>
     */
    protected $volume_path = '';
    /**
     * Generated from protobuf field <code>.models.fileio.QuotaEntry_6 default_quota = 2;</code>
     */
    protected $default_quota = null;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type string $volume_path
     *     @type \Models\Fileio\QuotaEntry_6 $default_quota
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Services\Fileio\DiskManagement::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>string volume_path = 1;</code>
     * @return string
     */
    public function getVolumePath()
    {
        return $this->volume_path;
    }

    /**
     * Generated from protobuf field <code>string volume_path = 1;</code>
     * @param string $var
     * @return $this
     */
    public function setVolumePath($var)
    {
        GPBUtil::checkString($var, True);
        $this->volume_path = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>.models.fileio.QuotaEntry_6 default_quota = 2;</code>
     * @return \Models\Fileio\QuotaEntry_6|null
     */
    public function getDefaultQuota()
    {
        return $this->default_quota;
    }

    public function hasDefaultQuota()
    {
        return isset($this->default_quota);
    }

    public function clearDefaultQuota()
    {
        unset($this->default_quota);
    }

    /**
     * Generated from protobuf field <code>.models.fileio.QuotaEntry_6 default_quota = 2;</code>
     * @param \Models\Fileio\QuotaEntry_6 $var
     * @return $this
     */
    public function setDefaultQuota($var)
    {
        GPBUtil::checkMessage($var, \Models\Fileio\QuotaEntry_6::class);
        $this->default_quota = $var;

        return $this;
    }

}
