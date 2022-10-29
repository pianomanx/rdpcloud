<?php
# Generated by the protocol buffer compiler.  DO NOT EDIT!
# source: models/msi/product.proto

namespace Models\Msi;

use Google\Protobuf\Internal\GPBType;
use Google\Protobuf\Internal\RepeatedField;
use Google\Protobuf\Internal\GPBUtil;

/**
 * Generated from protobuf message <code>models.msi.Product</code>
 */
class Product extends \Google\Protobuf\Internal\Message
{
    /**
     * Generated from protobuf field <code>string guid = 1;</code>
     */
    protected $guid = '';
    /**
     * Generated from protobuf field <code>string name = 2;</code>
     */
    protected $name = '';
    /**
     * Generated from protobuf field <code>string version = 3;</code>
     */
    protected $version = '';
    /**
     * Generated from protobuf field <code>string publisher = 4;</code>
     */
    protected $publisher = '';
    /**
     * Generated from protobuf field <code>.google.protobuf.Timestamp install_date = 5;</code>
     */
    protected $install_date = null;

    /**
     * Constructor.
     *
     * @param array $data {
     *     Optional. Data for populating the Message object.
     *
     *     @type string $guid
     *     @type string $name
     *     @type string $version
     *     @type string $publisher
     *     @type \Google\Protobuf\Timestamp $install_date
     * }
     */
    public function __construct($data = NULL) {
        \GPBMetadata\Models\Msi\Product::initOnce();
        parent::__construct($data);
    }

    /**
     * Generated from protobuf field <code>string guid = 1;</code>
     * @return string
     */
    public function getGuid()
    {
        return $this->guid;
    }

    /**
     * Generated from protobuf field <code>string guid = 1;</code>
     * @param string $var
     * @return $this
     */
    public function setGuid($var)
    {
        GPBUtil::checkString($var, True);
        $this->guid = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string name = 2;</code>
     * @return string
     */
    public function getName()
    {
        return $this->name;
    }

    /**
     * Generated from protobuf field <code>string name = 2;</code>
     * @param string $var
     * @return $this
     */
    public function setName($var)
    {
        GPBUtil::checkString($var, True);
        $this->name = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string version = 3;</code>
     * @return string
     */
    public function getVersion()
    {
        return $this->version;
    }

    /**
     * Generated from protobuf field <code>string version = 3;</code>
     * @param string $var
     * @return $this
     */
    public function setVersion($var)
    {
        GPBUtil::checkString($var, True);
        $this->version = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>string publisher = 4;</code>
     * @return string
     */
    public function getPublisher()
    {
        return $this->publisher;
    }

    /**
     * Generated from protobuf field <code>string publisher = 4;</code>
     * @param string $var
     * @return $this
     */
    public function setPublisher($var)
    {
        GPBUtil::checkString($var, True);
        $this->publisher = $var;

        return $this;
    }

    /**
     * Generated from protobuf field <code>.google.protobuf.Timestamp install_date = 5;</code>
     * @return \Google\Protobuf\Timestamp|null
     */
    public function getInstallDate()
    {
        return $this->install_date;
    }

    public function hasInstallDate()
    {
        return isset($this->install_date);
    }

    public function clearInstallDate()
    {
        unset($this->install_date);
    }

    /**
     * Generated from protobuf field <code>.google.protobuf.Timestamp install_date = 5;</code>
     * @param \Google\Protobuf\Timestamp $var
     * @return $this
     */
    public function setInstallDate($var)
    {
        GPBUtil::checkMessage($var, \Google\Protobuf\Timestamp::class);
        $this->install_date = $var;

        return $this;
    }

}

