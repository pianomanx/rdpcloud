<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: models/netmgmt/user.proto

namespace Models\Netmgmt;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>models.netmgmt.User_2</code>
 */
class User_2 extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>string password = 2;</code>
     */
    protected $password = '';

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type string $password
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Models\Netmgmt\User::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>string password = 2;</code>
     * @return string
     */
    public function getPassword()
    {
        return $this->password;
    }

    /**
     * Generated from protobuf field <code>string password = 2;</code>
     * @param string $var
     * @return $this
     */
    public function setPassword($var)
    {
        GPBUtil::checkString($var, True);
        $this->password = $var;

        return $this;
    }

}
