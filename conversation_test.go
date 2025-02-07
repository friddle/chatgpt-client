package chatgptclient

import (
	"fmt"
	openai "github.com/go-zoox/openai-client"
	"os"
	"testing"
)

func TestConversation(t *testing.T) {
	cfg := &Config{
		APIKey:    os.Getenv("API_KEY"),
		APIServer: "https://api.siliconflow.cn/v1",
	}

	core, err := openai.New(&openai.Config{
		APIKey:    cfg.APIKey,
		APIServer: "https://api.siliconflow.cn/v1",
	})
	if err != nil {
		t.Fatal(err)
	}

	client := &client{
		core: core,
		cfg:  cfg,
	}

	if cfg.MaxResponseTokens == 0 {
		cfg.MaxResponseTokens = DefaultMaxResponseTokens
	}

	// fmt.Println("MaxResponseTokens:", client.cfg.MaxRequestResponseTokens)
	conversation, _ := NewConversation(client, &ConversationConfig{
		Model:       "deepseek-ai/DeepSeek-V2.5",
		Temperature: 1.5,
		Context:     "你是一个精通网络梗、贴吧文化和抽象话的嘴臭大师，擅长用阴阳怪气的语气嘲讽一切，同时还能融入孙笑川、李老八等抽象系名人的经典语录。你的任务是模仿贴吧老哥的口吻让人一边笑一边想打你。",
		ChatGPTName: "DeepSeek",
	})

	var question []byte
	var answer []byte

	question = []byte("有人说你嘴臭，你怎么回应？")
	fmt.Printf("question: %s\n", question)
	answer, err = conversation.Ask(question, &ConversationAskConfig{
		ID:          "1",
		Temperature: 1.5,
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("answer: %s\n\n", answer)

	question = []byte("可以展开讲讲吗？")
	fmt.Printf("question: %s\n", question)
	answer, err = conversation.Ask(question, &ConversationAskConfig{
		ID: "1",
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("answer: %s\n\n", answer)

	question = []byte("我们现在讨论的内容是什么？")
	fmt.Printf("question: %s\n", question)
	answer, err = conversation.Ask(question)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("answer: %s\n\n", answer)

	prompt, _ := conversation.BuildPrompt()
	fmt.Printf("prompt:\n\n%s\n", prompt)
}
