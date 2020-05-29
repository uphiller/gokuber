package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/gin-gonic/gin"
	"net/http"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/v1/aws/")
	{
		v1.GET("/health", health)
		v1.POST("/cluster", setCluster)
	}
	return r
}

func main() {
	r := setupRouter()
	r.Run(":5002")
}

func health(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

func setCluster(c *gin.Context) {
	sess, err := session.NewSession(&aws.Config{
		Region:      aws.String("ap-northeast-2"),
		Credentials: credentials.NewStaticCredentials("AKIAS2KAI3ZLLOUPD7QQ", "An/sjce+59jTpuXEYDDqf8mxku1axAkPocJmReU7", ""),
	})

	svc := ec2.New(sess)

	// Specify the details of the instance that you want to create.
	runResult, err := svc.RunInstances(&ec2.RunInstancesInput{
		// An Amazon Linux AMI ID for t2.micro instances in the us-west-2 region
		ImageId:      aws.String("ami-01af223aa7f274198"),
		InstanceType: aws.String("t2.micro"),
		MinCount:     aws.Int64(1),
		MaxCount:     aws.Int64(1),
	})

	if err != nil {
		fmt.Println("Could not create instance", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"instance_id": *runResult.Instances[0].InstanceId,
		})
	}
}
