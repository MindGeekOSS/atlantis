package matchers

import (
	"reflect"

	"github.com/petergtz/pegomock"
	events "github.com/runatlantis/atlantis/server/events"
)

func AnyPtrToEventsCommentCommand() *events.CommentCommand {
	pegomock.RegisterMatcher(pegomock.NewAnyMatcher(reflect.TypeOf((*(*events.CommentCommand))(nil)).Elem()))
	var nullValue *events.CommentCommand
	return nullValue
}

func EqPtrToEventsCommentCommand(value *events.CommentCommand) *events.CommentCommand {
	pegomock.RegisterMatcher(&pegomock.EqMatcher{Value: value})
	var nullValue *events.CommentCommand
	return nullValue
}
