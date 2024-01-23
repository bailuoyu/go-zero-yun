package im

const SYNC_OTHER_MACHINE_Y = 1

const SYNC_OTHER_MACHINE_N = 2

const MSG_LIFE_TIME_DAY_7 = 604800

/** @var string 消息对象类型 - 自定义消息 */
const MSG_TYPE_CUSTOM = "TIMCustomElem"

/** @var string 自定义消息数据。 不作为 APNs 的 payload 字段下发，故从 payload 中无法获取 Data 字段。 */
const MSG_TYPE_CUSTOM_DATA = "Data"

/** @var string 自定义消息描述信息。当接收方为 iOS 或 Android 后台在线时，做离线推送文本展示。 */
const MSG_TYPE_CUSTOM_DESC = "Desc"

/** @var string 扩展字段。当接收方为 iOS 系统且应用处在后台时，此字段作为 APNs 请求包 Payloads 中的 Ext 键值下发，Ext 的协议格式由业务方确定，APNs 只做透传。 */
const MSG_TYPE_CUSTOM_EXT = "Ext"

/** @var string 自定义 APNs 推送铃音。 */
const MSG_TYPE_CUSTOM_SOUND = "Sound"

/**
 * 消息回调禁止开关，只对本条消息有效，ForbidBeforeSendMsgCallback 表示禁止发消息前回调，ForbidAfterSendMsgCallback 表示禁止发消息后回调
 * Array 选填
 */

const FORBID_BEFORE_SEND_MSG = "ForbidBeforeSendMsgCallback"
const FORBID_AFTER_SEND_MSG = "ForbidAfterSendMsgCallback"

/**
 * 消息发送控制选项，是一个 String 数组，只对本条消息有效。
 * "NoUnread"表示该条消息不计入未读数。
 * "NoLastMsg"表示该条消息不更新会话列表。
 * "WithMuteNotifications"表示该条消息的接收方对发送方设置的免打扰选项生效（默认不生效）。
 * 示例："SendMsgControl": ["NoUnread","NoLastMsg","WithMuteNotifications"]
 * Array 选填
 */
const SEND_MSG_CONTROL = "SendMsgControl"
const SEND_MSG_CONTROL_NOUNREAD = "NoUnread"
const SEND_MSG_CONTROL_NOLASTMSG = "NoLastMsg"
const SEND_MSG_CONTROL_WITH_MUTE = "WithMuteNotifications"
