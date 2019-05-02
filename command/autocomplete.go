package command

import (
	"github.com/minamijoyo/tfschema/tfschema"
	"github.com/posener/complete"
)

func (m *Meta) completePredictResourceType() complete.Predictor {
	return complete.PredictFunc(func(args complete.Args) []string {

		resourceType := args.Last
		providerName, err := detectProviderName(resourceType)
		if err != nil {
			return nil
		}

		client, err := tfschema.NewClient(providerName)
		if err != nil {
			return nil
		}

		defer client.Kill()

		resourceTypes, err := client.ResourceTypes()
		if err != nil {
			return nil
		}

		return resourceTypes

	})
}

func (m *Meta) completePredictDataSource() complete.Predictor {
	return complete.PredictFunc(func(args complete.Args) []string {

		dataSource := args.Last
		providerName, err := detectProviderName(dataSource)
		if err != nil {
			return nil
		}

		client, err := tfschema.NewClient(providerName)
		if err != nil {
			return nil
		}

		defer client.Kill()

		dataSources, err := client.DataSources()
		if err != nil {
			return nil
		}

		return dataSources
	})
}
