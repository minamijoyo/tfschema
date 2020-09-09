package command

import (
	"github.com/posener/complete"
)

func (m *Meta) completePredictResourceType() complete.Predictor {
	return complete.PredictFunc(func(args complete.Args) []string {

		resourceType := args.Last
		providerName, err := detectProviderName(resourceType)
		if err != nil {
			return nil
		}

		client, err := NewDefaultClient(providerName)
		if err != nil {
			return nil
		}

		defer client.Close()

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

		client, err := NewDefaultClient(providerName)
		if err != nil {
			return nil
		}

		defer client.Close()

		dataSources, err := client.DataSources()
		if err != nil {
			return nil
		}

		return dataSources
	})
}
