package gapi

import (
	"context"
	"github.com/haodam/Bank-Go/pb"
	"github.com/haodam/Bank-Go/val"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
)

func (server *Server) verifyEmail(ctx context.Context, req *pb.VerifyEmailRequest) (*pb.VerifyEmailResponse, error) {

	violations := validateVerifyEmailRequest(req)
	if violations != nil {
		return nil, invalidArgumentError(violations)
	}

	// converter data db.User on pb.User
	rsp := &pb.VerifyEmailResponse{}

	return rsp, nil
}

func validateVerifyEmailRequest(req *pb.VerifyEmailRequest) (violations []*errdetails.BadRequest_FieldViolation) {
	if err := val.ValidateEmailId(req.GetEmailId()); err != nil {
		violations = append(violations, fieldViolation("email_id", err))
	}
	if err := val.ValidateSecretCode(req.GetSecretCode()); err != nil {
		violations = append(violations, fieldViolation("secret_code", err))
	}
	return violations
}
