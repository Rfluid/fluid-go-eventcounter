package out

import (
	common_service "github.com/reb-felipe/eventcounter/cmd/common/service"
	"github.com/reb-felipe/eventcounter/cmd/config/flags"
	exchange_listener "github.com/reb-felipe/eventcounter/cmd/exchange-listener"
)

func GenerateConsumerOutputs() {
	// Gera os arquivos JSON com os resultados.
	common_service.CreateAndWriteFile(flags.OutputDir, "created", exchange_listener.Cw.CreatedCounts)
	common_service.CreateAndWriteFile(flags.OutputDir, "updated", exchange_listener.Cw.UpdatedCounts)
	common_service.CreateAndWriteFile(flags.OutputDir, "deleted", exchange_listener.Cw.DeletedCounts)
}
