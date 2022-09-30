<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: services/netmgmt/user.proto

namespace Services\Netmgmt;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>services.netmgmt.ChangeUserPasswordRequest</code>
 */
class ChangeUserPasswordRequest extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>.models.netmgmt.User_3 user = 1;</code>
     */
    protected $user = null;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type \Models\Netmgmt\User_3 $user
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Services\Netmgmt\User::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>.models.netmgmt.User_3 user = 1;</code>
     * @return \Models\Netmgmt\User_3|null
     */
    public function getUser()
    {
        return $this->user;
    }

    public function hasUser()
    {
        return isset($this->user);
    }

    public function clearUser()
    {
        unset($this->user);
    }

    /**
     * Generated from protobuf field <code>.models.netmgmt.User_3 user = 1;</code>
     * @param \Models\Netmgmt\User_3 $var
     * @return $this
     */
    public function setUser($var)
    {
        GPBUtil::checkMessage($var, \Models\Netmgmt\User_3::class);
        $this->user = $var;

        return $this;
    }

}

