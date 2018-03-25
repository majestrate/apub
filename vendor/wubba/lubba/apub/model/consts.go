package model

const ap_prefix = "http://activitystrea.ms/schema/1.0/"

const TypeNote = ap_prefix + "note"
const TypeActivity = ap_prefix + "activity"
const TypePerson = ap_prefix + "person"

const VerbPost = ap_prefix + "post"
const VerbShare = ap_prefix + "share"

const ActivityStreams = "http://www.w3.org/ns/activitystreams"
const W3Security = "https://w3id.org/security/v1"

var StdContext = []interface{}{
	ActivityStreams,
	W3Security,
	map[string]string{
		"toot":                      "http://joinmastodon.org/ns#",
		"sensative":                 "as:sensative",
		"ostatus":                   "http://ostatus.org#",
		"manuallyApprovesFollowers": "as:manuallyApprovesFollowers",
		"inReplyToAtomUri":          "ostatus:inReplyToAtomUri",
		"converstation":             "ostatus:conversation",
		"atomUri":                   "ostatus:atomUri",
		"HashTag":                   "as:Hashtag",
		"Emoji":                     "as:Emoji",
	},
}

const LinkType = "Link"
const ObjectType = "Object"
const PersonType = "Person"
const ImageType = "Image"
