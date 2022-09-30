<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: services/netmgmt/local_group.proto

namespace Services\Netmgmt;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>services.netmgmt.GetUsersInLocalGroupResponse</code>
 */
class GetUsersInLocalGroupResponse extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>repeated .models.netmgmt.User_1 users = 1;</code>
     */
    private $users;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type array<\Models\Netmgmt\User_1>|\Google\Protobuf\Internal\RepeatedField $users
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Services\Netmgmt\LocalGroup::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>repeated .models.netmgmt.User_1 users = 1;</code>
     * @return \Google\Protobuf\Internal\RepeatedField
     */
    public function getUsers()
    {
        return $this->users;
    }

    /**
     * Generated from protobuf field <code>repeated .models.netmgmt.User_1 users = 1;</code>
     * @param array<\Models\Netmgmt\User_1>|\Google\Protobuf\Internal\RepeatedField $var
     * @return $this
     */
    public function setUsers($var)
    {
        $arr = GPBUtil::checkRepeatedField($var, \Google\Protobuf\Internal\GPBType::MESSAGE, \Models\Netmgmt\User_1::class);
        $this->users = $arr;

        return $this;
    }

}

