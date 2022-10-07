<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: services/netmgmt/user.proto

namespace Services\Netmgmt;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>services.netmgmt.GetUserLocalGroupsResponse</code>
 */
class GetUserLocalGroupsResponse extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>repeated .models.netmgmt.LocalGroup_1 local_groups = 1;</code>
     */
    private $local_groups;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type array<\Models\Netmgmt\LocalGroup_1>|\Google\Protobuf\Internal\RepeatedField $local_groups
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Services\Netmgmt\User::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>repeated .models.netmgmt.LocalGroup_1 local_groups = 1;</code>
     * @return \Google\Protobuf\Internal\RepeatedField
     */
    public function getLocalGroups()
    {
        return $this->local_groups;
    }

    /**
     * Generated from protobuf field <code>repeated .models.netmgmt.LocalGroup_1 local_groups = 1;</code>
     * @param array<\Models\Netmgmt\LocalGroup_1>|\Google\Protobuf\Internal\RepeatedField $var
     * @return $this
     */
    public function setLocalGroups($var)
    {
        $arr = GPBUtil::checkRepeatedField($var, \Google\Protobuf\Internal\GPBType::MESSAGE, \Models\Netmgmt\LocalGroup_1::class);
        $this->local_groups = $arr;

        return $this;
    }

}
